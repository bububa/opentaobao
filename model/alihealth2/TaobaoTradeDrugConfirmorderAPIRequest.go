package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTradeDrugConfirmorderAPIRequest
阿里健康020接单 API请求
taobao.trade.drug.confirmorder

阿里健康020接单 */
type TaobaoTradeDrugConfirmorderAPIRequest struct {
	model.Params
	// 代送宝 代送商ID
	_deliveryId int64
	// public static int NORMAL_TYPE=0; 普通发货 默认 public static int DD_DAI_SONG=2; 代送宝	public static int DD_SONG_TYPE_V2=3; 点点送发货
	_confirmType int64
	// 订单ID
	_orderId int64
	// 子账号nick
	_subUserNick string
}

// New
