package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpmemberfindbrandinfolistAPIResponse 查询可用品牌列表 API返回值
// taobao.universalbp.member.findbrandinfolist
//
// 查询账号对应的品牌，用于品牌人群屏蔽等
type TaobaouniversalbpmemberfindbrandinfolistAPIResponse struct {
	model.CommonResponse
	TaobaouniversalbpmemberfindbrandinfolistAPIResponseModel
}

// TaobaouniversalbpmemberfindbrandinfolistAPIResponseModel is 查询可用品牌列表 成功返回结果
type TaobaouniversalbpmemberfindbrandinfolistAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_member_findbrandinfolist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaouniversalbpmemberfindbrandinfolistTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
