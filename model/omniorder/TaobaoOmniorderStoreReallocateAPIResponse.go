package omniorder

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoomniorderstorereallocateAPIResponse rellocate API返回值
// taobao.omniorder.store.reallocate
//
// 门店发货提供改派接口
type TaobaoomniorderstorereallocateAPIResponse struct {
	model.CommonResponse
	TaobaoomniorderstorereallocateAPIResponseModel
}

// TaobaoomniorderstorereallocateAPIResponseModel is rellocate 成功返回结果
type TaobaoomniorderstorereallocateAPIResponseModel struct {
	XMLName xml.Name `xml:"omniorder_store_reallocate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoomniorderstorereallocateResult `json:"result,omitempty" xml:"result,omitempty"`
}
