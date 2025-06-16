package portal

import (
	"math"

	"github.com/perfect-panel/server/internal/model/coupon"
	"github.com/perfect-panel/server/internal/model/payment"
	"github.com/perfect-panel/server/internal/types"
)

func getDiscount(discounts []types.SubscribeDiscount, inputMonths int64) float64 {
	var finalDiscount int64 = 100

	for _, discount := range discounts {
		if inputMonths >= discount.Quantity && discount.Discount < finalDiscount {
			finalDiscount = discount.Discount
		}
	}
	return float64(finalDiscount) / float64(100)
}

func calculateCoupon(amount int64, couponInfo *coupon.Coupon) int64 {
	if couponInfo.Type == 1 {
		return int64(float64(amount) * (float64(couponInfo.Discount) / float64(100)))
	} else {
		return min(couponInfo.Discount, amount)
	}
}

func calculateFee(amount int64, config *payment.Payment) int64 {
	var fee float64

	switch config.FeeMode {
	case 0:
		return 0

	case 1:
		// 固定金额的反推逻辑
		// 解方程：amount = x - 固定金额
		// 得到：x = amount + 固定金额
		total := float64(amount) / (1 - float64(config.FeePercent)/100)
		fee = total - float64(amount)

	case 2:
		if amount > 0 {
			fee = float64(config.FeeAmount)
		}

	case 3:
		// 百分比 + 固定手续费的反推逻辑
		// 解方程：amount = x - (x * 百分比 + 固定金额)
		// 得到：x = (amount + 固定金额) / (1 - 百分比)
		feeRate := float64(config.FeePercent) / 100
		total := (float64(amount) + float64(config.FeeAmount)) / (1 - feeRate)
		fee = total - float64(amount)

	}

	return int64(math.Round(fee)) // 四舍五入
}
