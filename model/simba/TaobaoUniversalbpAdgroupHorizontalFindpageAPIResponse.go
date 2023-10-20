package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpadgrouphorizontalfindpageAPIResponse 查询单元分页列表 API返回值
// taobao.universalbp.adgroup.horizontal.findpage
//
// 查询单元分页列表
type TaobaouniversalbpadgrouphorizontalfindpageAPIResponse struct {
	model.CommonResponse
	TaobaouniversalbpadgrouphorizontalfindpageAPIResponseModel
}

// TaobaouniversalbpadgrouphorizontalfindpageAPIResponseModel is 查询单元分页列表 成功返回结果
type TaobaouniversalbpadgrouphorizontalfindpageAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_adgroup_horizontal_findpage_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaouniversalbpadgrouphorizontalfindpageTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
