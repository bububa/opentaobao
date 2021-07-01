package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBaikeImportZhubaoDataAPIRequest
导入数据到商品百科服务 API请求
taobao.baike.import.zhubao.data

用于接入外部数据录入到商品百科中 */
type TaobaoBaikeImportZhubaoDataAPIRequest struct {
	model.Params
	// 约定的Json数据
	_dataJsonStr string
}

// New
