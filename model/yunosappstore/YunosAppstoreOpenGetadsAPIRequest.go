package yunosappstore

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosAppstoreOpenGetadsAPIRequest
获取外投广告 API请求
yunos.appstore.open.getads

将广告外投给外部合作伙伴 */
type YunosAppstoreOpenGetadsAPIRequest struct {
	model.Params
	// 请求id
	_rid string
	// 指定广告分类
	_cats []string
	// 是否排除已安装
	_excludeInstall bool
	// 场景或页面标识
	_caseId string
	// ssp标识
	_ssp string
	// 结算类型
	_feeType string
	// 客户端来源ip
	_clientIp string
	// 广告指定包名
	_pkgs []string
	// 客户端版本号
	_clientVerCode int64
	// 是否映射到uuid
	_tryMapToUuid bool
	// 排除包名列表
	_excludePkgs []string
	// 设备唯一标识
	_deviceId string
	// 广告数量
	_size int64
	// 排除分类
	_excludeCats []string
	// 创意模板id列表
	_templateIds []int64
	// 广告底价
	_mrp int64
	// 请求特征集
	_options int64
}

// New
