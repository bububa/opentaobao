package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpStdcategoryFindcategoryconditionAPIResponse 获取类目过滤条件 API返回值
// taobao.universalbp.stdcategory.findcategorycondition
//
// 查询全量类目信息列表，用于进行类目兴趣人群相关定向
type TaobaoUniversalbpStdcategoryFindcategoryconditionAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpStdcategoryFindcategoryconditionAPIResponseModel
}

// TaobaoUniversalbpStdcategoryFindcategoryconditionAPIResponseModel is 获取类目过滤条件 成功返回结果
type TaobaoUniversalbpStdcategoryFindcategoryconditionAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_stdcategory_findcategorycondition_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpStdcategoryFindcategoryconditionTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
