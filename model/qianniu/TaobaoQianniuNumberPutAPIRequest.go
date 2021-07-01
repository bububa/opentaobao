package qianniu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQianniuNumberPutAPIRequest
ISV上传数据接口 API请求
taobao.qianniu.number.put

ISV提供给卖家使用的业务数据，需要通过这个接口上传到千牛数据中心。 */
type TaobaoQianniuNumberPutAPIRequest struct {
	model.Params
	// 考虑到稳定性，建议一次卖家最多为200个。标准json格式的数组构成的字符串。每个元素为{user_id:****,field:"****",value:"****"}分别是用户的userid，数据的名称，以及数据的值。
	_data string
}

// New
