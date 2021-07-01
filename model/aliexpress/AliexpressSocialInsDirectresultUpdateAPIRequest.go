package aliexpress

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressSocialInsDirectresultUpdateAPIRequest
ISV更新INS私信发送的结果 API请求
aliexpress.social.ins.directresult.update

ISV更新INS私信发送的结果 */
type AliexpressSocialInsDirectresultUpdateAPIRequest struct {
	model.Params
	// 回调id,在获取图片时会返回
	_id int64
	// 接受消息人INS_ID，也就是查询图片时的request_ins_id
	_receiveInsId string
	// ISV发送私信人的INS_ID
	_senderInsId string
	// 1.成功，2.失败。
	_result int64
}

// New
