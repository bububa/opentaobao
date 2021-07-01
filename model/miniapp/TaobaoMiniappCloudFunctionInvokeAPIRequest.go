package miniapp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMiniappCloudFunctionInvokeAPIRequest
外部触发云函数 API请求
taobao.miniapp.cloud.function.invoke

用户isv从外部触发聚石塔云函数的执行。 */
type TaobaoMiniappCloudFunctionInvokeAPIRequest struct {
	model.Params
	// 云函数名称
	_name string
	// 指定云函数的handler
	_handler string
	// 调用云函数时的传参（JSON格式）
	_data string
	// 云环境
	_env string
	// 扩展协议参数
	_exts string
}

// New
