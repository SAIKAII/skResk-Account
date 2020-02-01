package accounts

import (
	"testing"

	"github.com/sirupsen/logrus"

	"github.com/SAIKAII/skResk-Account/services"

	"github.com/shopspring/decimal"

	"github.com/segmentio/ksuid"

	"github.com/SAIKAII/skResk-Infra/base"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/tietang/dbx"
)

func TestAccountLogDao(t *testing.T) {
	err := base.Tx(func(runner *dbx.TxRunner) error {
		dao := &AccountLogDao{
			runner: runner,
		}
		Convey("通过流水编号查询流水记录", t, func() {
			a := &AccountLog{
				LogNo:      ksuid.New().Next().String(),
				TradeNo:    ksuid.New().Next().String(),
				Status:     1,
				AccountNo:  ksuid.New().Next().String(),
				UserId:     ksuid.New().Next().String(),
				Username:   "测试用户",
				Amount:     decimal.NewFromFloat(1),
				Balance:    decimal.NewFromFloat(100),
				ChangeFlag: services.FlagAccountCreated,
				ChangeType: services.AccountCreated,
			}

			// 通过LogNo查询
			Convey("通过LogNo查询", func() {
				id, err := dao.Insert(a)
				So(err, ShouldBeNil)
				So(id, ShouldBeGreaterThan, 0)
				na := dao.GetOne(a.LogNo)
				So(na, ShouldNotBeNil)
				So(na.Balance, ShouldEqual, a.Balance)
				So(na.Amount.String(), ShouldEqual, a.Amount.String())
				So(na.CreatedAt, ShouldNotBeNil)
			})
			// 通过TradeNo查询
			Convey("通过TradeNo查询", func() {
				id, err := dao.Insert(a)
				So(err, ShouldBeNil)
				So(id, ShouldBeGreaterThan, 0)
				na := dao.GetByTradeNo(a.TradeNo)
				So(na, ShouldNotBeNil)
				So(na.Balance, ShouldEqual, a.Balance)
				So(na.Amount.String(), ShouldEqual, a.Amount.String())
				So(na.CreatedAt, ShouldNotBeNil)
			})
		})
		return nil
	})
	if err != nil {
		logrus.Error(err)
	}
}
