package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoProductImageUploadAPIResponse 产品图片上传 API返回值
// taobao.fenxiao.product.image.upload
//
// 产品主图图片空间相对路径或绝对路径添加或更新，或者是图片上传。如果指定位置的图片已存在，则覆盖原有信息。如果位置为1,自动设为主图；如果位置为0，表示属性图片
type TaobaoFenxiaoProductImageUploadAPIResponse struct {
	model.CommonResponse
	TaobaoFenxiaoProductImageUploadAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFenxiaoProductImageUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFenxiaoProductImageUploadAPIResponseModel).Reset()
}

// TaobaoFenxiaoProductImageUploadAPIResponseModel is 产品图片上传 成功返回结果
type TaobaoFenxiaoProductImageUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_product_image_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作时间
	Created string `json:"created,omitempty" xml:"created,omitempty"`
	// 操作是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFenxiaoProductImageUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Created = ""
	m.Result = false
}

var poolTaobaoFenxiaoProductImageUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFenxiaoProductImageUploadAPIResponse)
	},
}

// GetTaobaoFenxiaoProductImageUploadAPIResponse 从 sync.Pool 获取 TaobaoFenxiaoProductImageUploadAPIResponse
func GetTaobaoFenxiaoProductImageUploadAPIResponse() *TaobaoFenxiaoProductImageUploadAPIResponse {
	return poolTaobaoFenxiaoProductImageUploadAPIResponse.Get().(*TaobaoFenxiaoProductImageUploadAPIResponse)
}

// ReleaseTaobaoFenxiaoProductImageUploadAPIResponse 将 TaobaoFenxiaoProductImageUploadAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFenxiaoProductImageUploadAPIResponse(v *TaobaoFenxiaoProductImageUploadAPIResponse) {
	v.Reset()
	poolTaobaoFenxiaoProductImageUploadAPIResponse.Put(v)
}
