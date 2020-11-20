package websocket

import (
	"encoding/json"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/settings/webutils"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs"
	"github.com/iwind/TeaGo/actions"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "setting", "index")
	this.SecondMenu("websocket")
}

func (this *IndexAction) RunGet(params struct {
	ServerId int64
}) {
	webConfig, err := webutils.FindWebConfigWithServerId(this.Parent(), params.ServerId)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Data["webId"] = webConfig.Id
	this.Data["websocketRef"] = webConfig.WebsocketRef
	this.Data["websocketConfig"] = webConfig.Websocket

	this.Show()
}

func (this *IndexAction) RunPost(params struct {
	WebId            int64
	WebsocketRefJSON []byte
	WebsocketJSON    []byte

	Must *actions.Must
}) {
	defer this.CreateLogInfo("修改Web %d 的Websocket设置", params.WebId)

	// TODO 检查配置

	websocketRef := &serverconfigs.HTTPWebsocketRef{}
	err := json.Unmarshal(params.WebsocketRefJSON, websocketRef)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	websocketConfig := &serverconfigs.HTTPWebsocketConfig{}
	err = json.Unmarshal(params.WebsocketJSON, websocketConfig)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	err = websocketConfig.Init()
	if err != nil {
		this.ErrorPage(err)
		return
	}

	// 创建
	handshakeTimeoutJSON, err := json.Marshal(websocketConfig.HandshakeTimeout)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	// 创建或修改
	if websocketConfig.Id <= 0 {
		createResp, err := this.RPC().HTTPWebsocketRPC().CreateHTTPWebsocket(this.AdminContext(), &pb.CreateHTTPWebsocketRequest{
			HandshakeTimeoutJSON: handshakeTimeoutJSON,
			AllowAllOrigins:      websocketConfig.AllowAllOrigins,
			AllowedOrigins:       websocketConfig.AllowedOrigins,
			RequestSameOrigin:    websocketConfig.RequestSameOrigin,
			RequestOrigin:        websocketConfig.RequestOrigin,
		})
		if err != nil {
			this.ErrorPage(err)
			return
		}
		websocketConfig.Id = createResp.WebsocketId
	} else {
		_, err = this.RPC().HTTPWebsocketRPC().UpdateHTTPWebsocket(this.AdminContext(), &pb.UpdateHTTPWebsocketRequest{
			WebsocketId:          websocketConfig.Id,
			HandshakeTimeoutJSON: handshakeTimeoutJSON,
			AllowAllOrigins:      websocketConfig.AllowAllOrigins,
			AllowedOrigins:       websocketConfig.AllowedOrigins,
			RequestSameOrigin:    websocketConfig.RequestSameOrigin,
			RequestOrigin:        websocketConfig.RequestOrigin,
		})
	}

	websocketRef.WebsocketId = websocketConfig.Id
	websocketRefJSON, err := json.Marshal(websocketRef)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	_, err = this.RPC().HTTPWebRPC().UpdateHTTPWebWebsocket(this.AdminContext(), &pb.UpdateHTTPWebWebsocketRequest{
		WebId:         params.WebId,
		WebsocketJSON: websocketRefJSON,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
