package crm

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmGradeSetAPIRequest 卖家设置等级规则 API请求
// taobao.crm.grade.set
//
// 设置等级信息，可以设置层级等级，也可以单独设置一个等级。出于安全原因，折扣现最低只能设置到700即7折。
type TaobaoCrmGradeSetAPIRequest struct {
	model.Params
	// 只对设置的层级等级有效，必须要在amount和count参数中选择一个<br><br/>amount参数的填写规范：升级到下一个级别的需要的交易额，单位为分,必须全部填写.例如10000,20000,30000，其中10000表示非会员升级到普通的所需的交易额，20000表示普通升级到高级所需的交易额，层级等级中最高等级的下一个等级默认为0。会员等级越高，所需交易额必须越高。
	_amount []int64
	// 会员等级，用逗号分隔。买家会员级别0：店铺客户 1：普通会员 2 ：高级会员 3：VIP会员 4：至尊VIP
	_grade []int64
	// 只对设置的层级等级有效，必须要在amount和count参数中选择一个<br><br/>count参数的填写规范：<br/>升级到下一个级别的需要的交易量,必须全部填写. 以逗号分隔,例如100,200,300，其中100表示非会员升级到普通会员交易量。层级等级中最高等级的下一个等级的交易量默认为0。会员等级越高，交易量必须越高。
	_count []int64
	// 会员级别折扣率。会员等级越高，折扣必须越低。<br/>950即9.5折，888折即8.88折。出于安全原因，折扣现最低只能设置到700即7折。
	_discount []int64
	// 是否设置达到某一会员等级的交易量和交易额，必填。4个级别都需要设置，如入参为true,true,true,false时，表示设置达到高级会员、VIP会员的交易量或者交易额，不设置达到至尊会员的交易量和交易额
	_hierarchy []bool
}

// NewTaobaoCrmGradeSetRequest 初始化TaobaoCrmGradeSetAPIRequest对象
func NewTaobaoCrmGradeSetRequest() *TaobaoCrmGradeSetAPIRequest {
	return &TaobaoCrmGradeSetAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoCrmGradeSetAPIRequest) Reset() {
	r._amount = r._amount[:0]
	r._grade = r._grade[:0]
	r._count = r._count[:0]
	r._discount = r._discount[:0]
	r._hierarchy = r._hierarchy[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCrmGradeSetAPIRequest) GetApiMethodName() string {
	return "taobao.crm.grade.set"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCrmGradeSetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoCrmGradeSetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAmount is Amount Setter
// 只对设置的层级等级有效，必须要在amount和count参数中选择一个&lt;br&gt;&lt;br/&gt;amount参数的填写规范：升级到下一个级别的需要的交易额，单位为分,必须全部填写.例如10000,20000,30000，其中10000表示非会员升级到普通的所需的交易额，20000表示普通升级到高级所需的交易额，层级等级中最高等级的下一个等级默认为0。会员等级越高，所需交易额必须越高。
func (r *TaobaoCrmGradeSetAPIRequest) SetAmount(_amount []int64) error {
	r._amount = _amount
	r.Set("amount", _amount)
	return nil
}

// GetAmount Amount Getter
func (r TaobaoCrmGradeSetAPIRequest) GetAmount() []int64 {
	return r._amount
}

// SetGrade is Grade Setter
// 会员等级，用逗号分隔。买家会员级别0：店铺客户 1：普通会员 2 ：高级会员 3：VIP会员 4：至尊VIP
func (r *TaobaoCrmGradeSetAPIRequest) SetGrade(_grade []int64) error {
	r._grade = _grade
	r.Set("grade", _grade)
	return nil
}

// GetGrade Grade Getter
func (r TaobaoCrmGradeSetAPIRequest) GetGrade() []int64 {
	return r._grade
}

// SetCount is Count Setter
// 只对设置的层级等级有效，必须要在amount和count参数中选择一个&lt;br&gt;&lt;br/&gt;count参数的填写规范：&lt;br/&gt;升级到下一个级别的需要的交易量,必须全部填写. 以逗号分隔,例如100,200,300，其中100表示非会员升级到普通会员交易量。层级等级中最高等级的下一个等级的交易量默认为0。会员等级越高，交易量必须越高。
func (r *TaobaoCrmGradeSetAPIRequest) SetCount(_count []int64) error {
	r._count = _count
	r.Set("count", _count)
	return nil
}

// GetCount Count Getter
func (r TaobaoCrmGradeSetAPIRequest) GetCount() []int64 {
	return r._count
}

// SetDiscount is Discount Setter
// 会员级别折扣率。会员等级越高，折扣必须越低。&lt;br/&gt;950即9.5折，888折即8.88折。出于安全原因，折扣现最低只能设置到700即7折。
func (r *TaobaoCrmGradeSetAPIRequest) SetDiscount(_discount []int64) error {
	r._discount = _discount
	r.Set("discount", _discount)
	return nil
}

// GetDiscount Discount Getter
func (r TaobaoCrmGradeSetAPIRequest) GetDiscount() []int64 {
	return r._discount
}

// SetHierarchy is Hierarchy Setter
// 是否设置达到某一会员等级的交易量和交易额，必填。4个级别都需要设置，如入参为true,true,true,false时，表示设置达到高级会员、VIP会员的交易量或者交易额，不设置达到至尊会员的交易量和交易额
func (r *TaobaoCrmGradeSetAPIRequest) SetHierarchy(_hierarchy []bool) error {
	r._hierarchy = _hierarchy
	r.Set("hierarchy", _hierarchy)
	return nil
}

// GetHierarchy Hierarchy Getter
func (r TaobaoCrmGradeSetAPIRequest) GetHierarchy() []bool {
	return r._hierarchy
}

var poolTaobaoCrmGradeSetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoCrmGradeSetRequest()
	},
}

// GetTaobaoCrmGradeSetRequest 从 sync.Pool 获取 TaobaoCrmGradeSetAPIRequest
func GetTaobaoCrmGradeSetAPIRequest() *TaobaoCrmGradeSetAPIRequest {
	return poolTaobaoCrmGradeSetAPIRequest.Get().(*TaobaoCrmGradeSetAPIRequest)
}

// ReleaseTaobaoCrmGradeSetAPIRequest 将 TaobaoCrmGradeSetAPIRequest 放入 sync.Pool
func ReleaseTaobaoCrmGradeSetAPIRequest(v *TaobaoCrmGradeSetAPIRequest) {
	v.Reset()
	poolTaobaoCrmGradeSetAPIRequest.Put(v)
}
