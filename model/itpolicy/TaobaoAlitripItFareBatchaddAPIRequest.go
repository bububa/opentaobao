package itpolicy

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripItFareBatchaddAPIRequest
【国际机票自有政策】批量添加 API请求
taobao.alitrip.it.fare.batchadd

支持自有政策和销售规则批量添加，支持携程的数据格式。淘宝格式为list [object] to json string，object的属性和单条接口一致。每个接入方最多同时只能有1个处理中的导入任务，超过后直接返回失败。文件一定要zip压缩，压缩后大小不超过5M，编码格式utf-8 */
type TaobaoAlitripItFareBatchaddAPIRequest struct {
	model.Params
	// 新增类型，1 自有政策单程 2 自有政策往返 3 销售规则
	_addType int64
	// 文本zip压缩后的数据字节流
	_bytes *model.File
	// 数据格式类型，1 淘宝 2 携程
	_dataType int64
	// json格式的字符串，扩展属性，预留
	_extendAttributes string
}

// New
