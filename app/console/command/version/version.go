package version

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zhaoyang1214/ginco/framework/contract"
	"os"
	"time"
)

func Command(a contract.Application) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Get Application version",
		Long:  "Get Application version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("The Application version is v%s\n", a.Version())
			c := a.GetIgnore("cache").(contract.Cache)
			ctx := context.Background()

			v, err := c.Get(ctx, "test2")
			fmt.Println(err)
			if err != nil {
				fmt.Println(string(v))
			}

			err = c.Set(ctx, "test2", []byte("2222"), time.Minute)
			err = c.Set(ctx, "te3", []byte("333"), time.Minute)
			err = c.Set(ctx, "te4", []byte("444"), time.Minute)
			fmt.Println(err)
			v, err = c.Get(ctx, "test2")
			_ = c.Delete(ctx, "test1", "test2")
			c.Has(ctx, "test")
			fmt.Println(v, err)
			_ = c.Clear(ctx)
			//fmt.Println(c.ClearPrefix(ctx, "te"))
			v, err = c.Get(ctx, "test2")
			fmt.Println(v, err)
			v, err = c.Get(ctx, "te3")
			fmt.Println(v, err)
			os.Exit(0)

			/*	_ = c.Set(ctx, "test1", []byte("111"), time.Minute)

				v, _ := c.Get(ctx, "test1")
				println(string(v))
				println(c.Has(ctx, "test1"))
				println(c.Has(ctx, "test2"))

				_ = c.Set(ctx, "t2", []byte("22"), time.Minute)
				_ = c.Set(ctx, "t3", []byte("33"), time.Minute)

				v, _ = c.Get(ctx, "t2")
				println(string(v))

				println(c.Has(ctx, "t3"))
				_ = c.Delete(ctx, "t3")
				println(c.Has(ctx, "t3"))*/

			//_ = c.Clear(ctx)

		},
	}
}
