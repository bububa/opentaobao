package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMixnickChangeAPIRequest
新旧mixnick互转 API请求
taobao.mixnick.change

如果采用老的Appkey获取MixNick A’， 发现DB里 new MixNick为null，则使用新的Appkey调API来换取A’’；如果采用新的Appkey获取A’’，发现DB里不存在当前A’’ 的记录时，先用老Appkey调API得到A’ 查询是否存在A用户的记录，如果已经存在，则关联，否则新增一条MixNick为null的新用户记录。 */
type TaobaoMixnickChangeAPIRequest struct {
	model.Params
	// 输入的混淆nick
	_srcMixNick string
	// 输入的appkey
	_srcAppkey string
}

// New
