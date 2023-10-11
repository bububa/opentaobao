package simba

// AdgroupVo 结构体
type AdgroupVo struct {
	// 计划名称
	CampaignName string `json:"campaign_name,omitempty" xml:"campaign_name,omitempty"`
	// 单元名称
	AdgroupName string `json:"adgroup_name,omitempty" xml:"adgroup_name,omitempty"`
	// 前端展示状态,pause:暂停推广,start:正在推广,terminate:结束推广,wait:等待推广,pending:准备推广,wait_pay:计划未付款
	DisplayStatus string `json:"display_status,omitempty" xml:"display_status,omitempty"`
	// 审核拒绝原因
	AuditReason string `json:"audit_reason,omitempty" xml:"audit_reason,omitempty"`
	// 推广主体类型,item:商品,item_private_mini:独享橱窗,shop:店铺,content:内容,short_video:短视频,user_define:自定义;
	PromotionType string `json:"promotion_type,omitempty" xml:"promotion_type,omitempty"`
	// 推广子主体类型,item:商品，item_private_mini:独享橱窗，shop:店铺，you_hao_huo:有好货，picture:图文，short_video:短视频，live_room:直播间，live_spot:看点，tao_blocks:淘积木，user_define_url:自定义url
	SubPromotionType string `json:"sub_promotion_type,omitempty" xml:"sub_promotion_type,omitempty"`
	// 视频组件名称
	AliveGroupName string `json:"alive_group_name,omitempty" xml:"alive_group_name,omitempty"`
	// 智能创意生成状态,success:合成成功,fail:合成失败
	BlackCreativeStatus string `json:"black_creative_status,omitempty" xml:"black_creative_status,omitempty"`
	// 计划id,计划已经存在场景必填,eg:添加主体、编辑计划状态等场景
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
	// 单元id
	AdgroupId int64 `json:"adgroup_id,omitempty" xml:"adgroup_id,omitempty"`
	// 投放状态,-1:未进入投放状态,0:用户设置状态-下线,1:用户设置状态-上线,2:合同未生效,9:投放结束,105:待支付
	OnlineStatus int64 `json:"online_status,omitempty" xml:"online_status,omitempty"`
	// 风控审核状态,-1:待审核,0:后台下线状态,1:后台上线状态
	AuditStatus int64 `json:"audit_status,omitempty" xml:"audit_status,omitempty"`
	// 单元智能出价信息
	AdgroupOcpc *AdgroupOcpcVo `json:"adgroup_ocpc,omitempty" xml:"adgroup_ocpc,omitempty"`
	// 主体类型对应的物料信息
	Material *CommonMaterialVo `json:"material,omitempty" xml:"material,omitempty"`
	// 视频组件id,0:宝贝链接,3:订阅店铺,5:直播,6:加购,7:收藏加购,8:入会有礼,12:关注有礼,11:直播联动,10:自定义模板,13:同店搜
	AliveGroupId int64 `json:"alive_group_id,omitempty" xml:"alive_group_id,omitempty"`
}
