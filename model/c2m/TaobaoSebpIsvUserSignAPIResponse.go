package c2m

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSebpIsvUserSignAPIResponse
淘小铺三方签约同步 API返回值
taobao.sebp.isv.user.sign

同步淘小铺三方服务签约信息 */
type TaobaoSebpIsvUserSignAPIResponse struct {
	model.CommonResponse
	TaobaoSebpIsvUserSignAPIResponseModel
}

// TaobaoSebpIsvUserSignAPIResponseModel is 淘小铺三方签约同步 成功返回结果
type TaobaoSebpIsvUserSignAPIResponseModel struct {
	XMLName xml.Name `xml:"sebp_isv_user_sign_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
