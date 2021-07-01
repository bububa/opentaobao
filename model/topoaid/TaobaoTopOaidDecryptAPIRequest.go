package topoaid

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTopOaidDecryptAPIRequest
OAID解密 API请求
taobao.top.oaid.decrypt

解码OAID(Open Addressee ID)，返回收件人信息。 */
type TaobaoTopOaidDecryptAPIRequest struct {
	model.Params
	// 解密请求列表，最多支持20个。
	_queryList []ReceiverQuery
}

// New
