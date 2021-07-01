package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJstAstrolabeStoreinventoryIteminitialAPIRequest
库存初始化接口 API请求
taobao.jst.astrolabe.storeinventory.iteminitial

ERP调用奇门的接口，对门店的库存进行初始化 */
type TaobaoJstAstrolabeStoreinventoryIteminitialAPIRequest struct {
	model.Params
	// 门店列表
	_stores []Store
	// 操作时间
	_operationTime string
}

// New
