package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpshopcategoryfindlistAPIResponse 人群相关类目查询 API返回值
// taobao.universalbp.shopcategory.findlist
//
// 查询店铺所属的类目信息
type TaobaouniversalbpshopcategoryfindlistAPIResponse struct {
	model.CommonResponse
	TaobaouniversalbpshopcategoryfindlistAPIResponseModel
}

// TaobaouniversalbpshopcategoryfindlistAPIResponseModel is 人群相关类目查询 成功返回结果
type TaobaouniversalbpshopcategoryfindlistAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_shopcategory_findlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaouniversalbpshopcategoryfindlistTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
