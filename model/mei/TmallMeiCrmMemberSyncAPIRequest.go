package mei

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallMeiCrmMemberSyncAPIRequest
同步推送会员信息 API请求
tmall.mei.crm.member.sync

品牌方通过该api主动推送会员信息。使用场景包括
1.用户在线上注册后，线下补充信息后，同步到线上。
2.其他情况的主动推送变更。 */
type TmallMeiCrmMemberSyncAPIRequest struct {
	model.Params
	// 会员手机号码
	_mobile string
	// 会员积分
	_point int64
	// 会员等级
	_level int64
	// 会员拓展信息
	_extend string
	// 该次同步的版本信息（建议使用时间戳）
	_version int64
	// 混淆昵称
	_mixNick string
	// 昵称
	_nick string
}

// New
