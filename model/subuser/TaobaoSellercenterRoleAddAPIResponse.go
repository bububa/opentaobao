package subuser

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSellercenterRoleAddAPIResponse 子账号角色的新增（指定卖家） API返回值
// taobao.sellercenter.role.add
//
// 给指定的卖家创建新的子账号角色&lt;br/&gt;&lt;br/&gt;如果需要授权的权限点有下级权限点或上级权限点，把该权限点的父权限点和该权限点的所有子权限都一并做赋权操作，并递归处理&lt;br/&gt;例如：权限点列表如下&lt;br/&gt;&lt;br/&gt;code=sell 宝贝管理&lt;br/&gt;&lt;br/&gt;---------|code=sm 店铺管理&lt;br/&gt;&lt;br/&gt;---------|---------|code=sm-design 如店铺装修&lt;br/&gt;&lt;br/&gt;---------|---------|---------|code=sm-tbd-visit内店装修入口&lt;br/&gt;&lt;br/&gt;---------|---------|---------|code=sm-tbd-publish内店装修发布&lt;br/&gt;&lt;br/&gt;---------|---------|code=phone 手机淘宝店铺&lt;br/&gt;&lt;br/&gt;调用改接口给code=sm-design店铺装修赋权时，同时会将下列权限点都赋予默认角色&lt;br/&gt;&lt;br/&gt;code=sell 宝贝管理&lt;br/&gt;&lt;br/&gt;---------|code=sm 店铺管理&lt;br/&gt;&lt;br/&gt;---------|---------|code=sm-design 如店铺装修&lt;br/&gt;&lt;br/&gt;---------|---------|---------|code=sm-tbd-visit内店装修入口&lt;br/&gt;&lt;br/&gt;---------|---------|---------|code=sm-tbd-publish内店装修发布&lt;br/&gt;
type TaobaoSellercenterRoleAddAPIResponse struct {
	model.CommonResponse
	TaobaoSellercenterRoleAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSellercenterRoleAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSellercenterRoleAddAPIResponseModel).Reset()
}

// TaobaoSellercenterRoleAddAPIResponseModel is 子账号角色的新增（指定卖家） 成功返回结果
type TaobaoSellercenterRoleAddAPIResponseModel struct {
	XMLName xml.Name `xml:"sellercenter_role_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 子账号角色
	Role *Role `json:"role,omitempty" xml:"role,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSellercenterRoleAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Role = nil
}

var poolTaobaoSellercenterRoleAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSellercenterRoleAddAPIResponse)
	},
}

// GetTaobaoSellercenterRoleAddAPIResponse 从 sync.Pool 获取 TaobaoSellercenterRoleAddAPIResponse
func GetTaobaoSellercenterRoleAddAPIResponse() *TaobaoSellercenterRoleAddAPIResponse {
	return poolTaobaoSellercenterRoleAddAPIResponse.Get().(*TaobaoSellercenterRoleAddAPIResponse)
}

// ReleaseTaobaoSellercenterRoleAddAPIResponse 将 TaobaoSellercenterRoleAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSellercenterRoleAddAPIResponse(v *TaobaoSellercenterRoleAddAPIResponse) {
	v.Reset()
	poolTaobaoSellercenterRoleAddAPIResponse.Put(v)
}
