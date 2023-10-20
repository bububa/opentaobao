package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpstdcategoryfindlistAPIResponse 人群相关类目查询 API返回值
// taobao.universalbp.stdcategory.findlist
//
// 入参商品信息，出参商品所属类别
type TaobaouniversalbpstdcategoryfindlistAPIResponse struct {
	model.CommonResponse
	TaobaouniversalbpstdcategoryfindlistAPIResponseModel
}

// TaobaouniversalbpstdcategoryfindlistAPIResponseModel is 人群相关类目查询 成功返回结果
type TaobaouniversalbpstdcategoryfindlistAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_stdcategory_findlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaouniversalbpstdcategoryfindlistTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
