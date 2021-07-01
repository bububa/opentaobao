package c2m

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSebpIsvUserCheckAPIResponse
淘小铺账户实名校验接口 API返回值
taobao.sebp.isv.user.check

校验淘小铺账户和身份信息匹配成功 */
type TaobaoSebpIsvUserCheckAPIResponse struct {
	model.CommonResponse
	TaobaoSebpIsvUserCheckAPIResponseModel
}

// TaobaoSebpIsvUserCheckAPIResponseModel is 淘小铺账户实名校验接口 成功返回结果
type TaobaoSebpIsvUserCheckAPIResponseModel struct {
	XMLName xml.Name `xml:"sebp_isv_user_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
