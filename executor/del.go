package executor

import (
	"context"
	"pegic/executor/util"
	"time"

	"github.com/XiaoMi/pegasus-go-client/pegasus"
)

func Del(rootCtx *Context, tb pegasus.TableConnector, hashKey *util.PegicBytes, sortkey *util.PegicBytes) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return tb.Del(ctx, hashKey.Bytes(), sortkey.Bytes())
}
