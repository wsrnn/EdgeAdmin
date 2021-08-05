package web

import (
	_ "github.com/TeaOSLab/EdgeAdmin/internal/tasks"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/about"

	// 系统用户
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/admins"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/admins/recipients"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/admins/recipients/groups"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/admins/recipients/instances"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/admins/recipients/logs"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/admins/recipients/tasks"

	// API节点
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/api"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/api/node"

	// 节点集群
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/clusters"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/clusters/cluster"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/clusters/cluster/settings"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/clusters/grants"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/clusters/logs"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/clusters/regions"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/clusters/regions/items"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/clusters/tasks"

	// 通用
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/csrf"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/dashboard"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/db"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/dns"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/dns/tasks"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/finance"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/finance/bills"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/index"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/log"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/logout"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/messages"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/nodes"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/ui"

	// 域名服务
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/ns"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/ns/clusters"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/ns/clusters/accessLogs"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/ns/clusters/cluster"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/ns/clusters/cluster/settings"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/ns/clusters/cluster/settings/accessLog"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/ns/clusters/logs"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/ns/routes"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/ns/settings/accesslogs"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/ns/users"

	// 服务相关
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/accesslogs"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/certs"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/components"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/components/cache"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/components/groups"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/components/log"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/components/waf"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/metrics"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/metrics/charts"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/boards"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/delete"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/log"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/settings"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/settings/access"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/settings/accessLog"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/settings/cache"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/settings/charset"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/settings/conds"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/settings/dns"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/settings/fastcgi"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/settings/gzip"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/settings/headers"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/settings/http"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/settings/https"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/settings/locations"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/settings/locations/access"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/settings/locations/accessLog"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/settings/locations/cache"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/settings/locations/charset"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/settings/locations/fastcgi"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/settings/locations/gzip"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/settings/locations/headers"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/settings/locations/http"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/settings/locations/location"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/settings/locations/pages"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/settings/locations/reverseProxy"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/settings/locations/rewrite"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/settings/locations/stat"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/settings/locations/waf"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/settings/locations/web"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/settings/locations/websocket"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/settings/origins"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/settings/pages"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/settings/redirects"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/settings/reverseProxy"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/settings/rewrite"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/settings/serverNames"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/settings/stat"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/settings/tcp"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/settings/tls"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/settings/udp"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/settings/unix"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/settings/waf"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/settings/web"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/settings/websocket"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/server/stat"

	// IP相关
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/ipbox"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/servers/iplists"

	// 设置相关
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/settings"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/settings/authority"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/settings/authority/nodes"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/settings/authority/nodes/node"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/settings/backup"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/settings/database"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/settings/ip-library"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/settings/login"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/settings/monitor-nodes"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/settings/monitor-nodes/node"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/settings/profile"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/settings/security"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/settings/server"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/settings/ui"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/settings/upgrade"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/settings/user-nodes"
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/settings/user-ui"

	// 恢复
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/recover"

	// 安装
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/setup"

	// 平台用户
	_ "github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/users"
)
