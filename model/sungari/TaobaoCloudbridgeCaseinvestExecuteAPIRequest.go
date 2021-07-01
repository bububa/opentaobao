package sungari

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoCloudbridgeCaseinvestExecuteAPIRequest
红盾云桥案件协查服务 API请求
taobao.cloudbridge.caseinvest.execute

通过API接口直接提供政府部门录入及查询函件服务 */
type TaobaoCloudbridgeCaseinvestExecuteAPIRequest struct {
	model.Params
	// 方法名称
	_apiName string
	// 方法参数
	_data string
}

// New
