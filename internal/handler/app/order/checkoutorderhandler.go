package order

import (
	"github.com/gin-gonic/gin"
	"github.com/perfect-panel/server/internal/logic/app/order"
	"github.com/perfect-panel/server/internal/svc"
	"github.com/perfect-panel/server/internal/types"
	"github.com/perfect-panel/server/pkg/result"
)

// Checkout order
func CheckoutOrderHandler(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req types.CheckoutOrderRequest
		_ = c.ShouldBind(&req)
		validateErr := svcCtx.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(c, validateErr)
			return
		}

		l := order.NewCheckoutOrderLogic(c.Request.Context(), svcCtx)
		resp, err := l.CheckoutOrder(&req, c.Request.Host)
		result.HttpResult(c, resp, err)
	}
}
