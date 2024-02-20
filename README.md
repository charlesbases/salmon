# Salmon

[[ants]](github.com/panjf2000/ants/v2)

```
golang 协程池, 对 `github.com/panjf2000/ants/v2` 的二次封装
```

------

## 修改部分

- 代码内封装了 sync.WaitGroup.Add() 和 sync.WaitGroup.Done() 、sync.WaitGroup.Wait() 函数的自动调用，

  外部调用只需要关注业务逻辑

- 添加 cancel 函数，调用后，阻塞中的任务将不执行
