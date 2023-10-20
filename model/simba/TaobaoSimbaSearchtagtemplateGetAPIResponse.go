package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbasearchtagtemplategetAPIResponse 获取搜索人群TOP用户可添加人群信息 API返回值
// taobao.simba.searchtagtemplate.get
//
// 获取搜索人群用户可添加人群信息
type TaobaosimbasearchtagtemplategetAPIResponse struct {
	model.CommonResponse
	TaobaosimbasearchtagtemplategetAPIResponseModel
}

// TaobaosimbasearchtagtemplategetAPIResponseModel is 获取搜索人群TOP用户可添加人群信息 成功返回结果
type TaobaosimbasearchtagtemplategetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_searchtagtemplate_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	TemplateList []TaobaosimbasearchtagtemplategetResult `json:"template_list,omitempty" xml:"template_list>taobaosimbasearchtagtemplateget_result,omitempty"`
}
