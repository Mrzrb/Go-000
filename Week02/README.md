学习笔记

1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

我认为遇到 `err == sql.ErrNoRows` 的时候不应该Wrap这个error, 因为dao层本来就是处理curd相关的逻辑, 如果sql.ErrNoRows的时候应该在内部处理 直接返回空数组和nil, 当遇到非dao逻辑的意外错误的时候 应该wrap 再向上抛
