package paimai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPaimaiItempropsGetAPIRequest
拍卖相关类目属性 API请求
taobao.paimai.itemprops.get

读取拍卖相关类目属性 */
type TaobaoPaimaiItempropsGetAPIRequest struct {
	model.Params
	// 获取类目的类型：1代表集市、2代表天猫
	_type int64
	// 是否商品属性，这个属性只能放于发布商品时使用。可选值:true(是),false(否)
	_isItemProp bool
	// 在is_enum_prop是true的前提下，是否是卖家可以自行输入的属性（注：如果is_enum_prop返回false，该参数统一返回false）。可选值:true(是),false(否) (删除的属性不会匹配和返回这个条件)
	_isInputProp bool
	// 是否枚举属性。可选值:true(是),false(否) (删除的属性不会匹配和返回这个条件)。如果返回true，属性值是下拉框选择输入，如果返回false，属性值是用户自行手工输入。
	_isEnumProp bool
	// 属性id (取类目属性时，传pid，不用同时传PID和parent_pid)
	_pid int64
	// 叶子类目ID，如果只传cid，则只返回一级属性,通过taobao.itemcats.get获得叶子类目ID
	_cid int64
	// 属性的Key，支持多条，以“,”分隔
	_attrKeys []string
	// 类目子属性路径,由该子属性上层的类目属性和类目属性值组成,格式pid:vid;pid:vid.取类目子属性需要传child_path,cid
	_childPath string
	// 父属性ID
	_parentPid int64
	// 是否销售属性。可选值:true(是),false(否)
	_isSaleProp bool
	// 增量时间戳。格式:yyyy-MM-dd HH:mm:ss假如传2005-01-01 00:00:00，则取所有的属性和子属性ID(如果传了pid会忽略datetime)
	_datetime string
	// 是否颜色属性。可选值:true(是),false(否) (删除的属性不会匹配和返回这个条件)
	_isColorProp bool
	// 是否关键属性。可选值:true(是),false(否)
	_isKeyProp bool
	// 需要返回的字段列表，见：ItemProp，默认返回：pid, name, must, multi, prop_values
	_fields []string
}

// New
