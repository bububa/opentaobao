package subuser

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/subuser"
)

// TaobaoSellercenterRoleAdd 子账号角色的新增（指定卖家）
// taobao.sellercenter.role.add
//
// 给指定的卖家创建新的子账号角色&lt;br/&gt;&lt;br/&gt;如果需要授权的权限点有下级权限点或上级权限点，把该权限点的父权限点和该权限点的所有子权限都一并做赋权操作，并递归处理&lt;br/&gt;例如：权限点列表如下&lt;br/&gt;&lt;br/&gt;code=sell 宝贝管理&lt;br/&gt;&lt;br/&gt;---------|code=sm 店铺管理&lt;br/&gt;&lt;br/&gt;---------|---------|code=sm-design 如店铺装修&lt;br/&gt;&lt;br/&gt;---------|---------|---------|code=sm-tbd-visit内店装修入口&lt;br/&gt;&lt;br/&gt;---------|---------|---------|code=sm-tbd-publish内店装修发布&lt;br/&gt;&lt;br/&gt;---------|---------|code=phone 手机淘宝店铺&lt;br/&gt;&lt;br/&gt;调用改接口给code=sm-design店铺装修赋权时，同时会将下列权限点都赋予默认角色&lt;br/&gt;&lt;br/&gt;code=sell 宝贝管理&lt;br/&gt;&lt;br/&gt;---------|code=sm 店铺管理&lt;br/&gt;&lt;br/&gt;---------|---------|code=sm-design 如店铺装修&lt;br/&gt;&lt;br/&gt;---------|---------|---------|code=sm-tbd-visit内店装修入口&lt;br/&gt;&lt;br/&gt;---------|---------|---------|code=sm-tbd-publish内店装修发布&lt;br/&gt;
func TaobaoSellercenterRoleAdd(ctx context.Context, clt *core.SDKClient, req *subuser.TaobaoSellercenterRoleAddAPIRequest, resp *subuser.TaobaoSellercenterRoleAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
