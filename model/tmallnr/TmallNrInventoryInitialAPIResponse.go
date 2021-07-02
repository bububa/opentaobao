package tmallnr

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallNrInventoryInitialAPIResponse 门店库存初始化前后端商品绑定 API返回值
// tmall.nr.inventory.initial
//
// 用于门店业务的商品的初始化，前端商品和后端商品绑定，走后端库存模式
type TmallNrInventoryInitialAPIResponse struct {
	model.CommonResponse
	TmallNrInventoryInitialAPIResponseModel
}

// TmallNrInventoryInitialAPIResponseModel is 门店库存初始化前后端商品绑定 成功返回结果
type TmallNrInventoryInitialAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nr_inventory_initial_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 初始化成功
	ResultData bool `json:"result_data,omitempty" xml:"result_data,omitempty"`
	// 错误信息提示
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误编码
	ErrorCode2 string `json:"error_code2,omitempty" xml:"error_code2,omitempty"`
	// isSuccess
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
