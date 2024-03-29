**创建一个 Web 服务器，用户可以在其中跟踪玩家赢了多少场游戏。**
https://studygolang.gitbook.io/learn-go-with-tests/gou-jian-ying-yong-cheng-xu/http-server

- `GET /players/{name}` 应该返回一个表示获胜总数的数字
- `POST /players/{name}` 应该为玩家赢得游戏记录一次得分，并随着每次 `POST` 递增

将遵循 TDD 方法，尽可能快地让程序先可用，然后进行小步迭代改进，直到我们找到解决方案。通过采取这种方法我们
- 在任何给定时间保持问题都是小问题
- 不要陷入陷阱(rabbit holes)
- 如果我们卡住或迷失了方向，回退不会前功尽弃

### 持续迭代，重构
我们强调了编写测试并**观察失败（红色）**，编写 _最少量_ 代码跑通测试（绿色）然后重构的 TDD 过程。

#### 先写测试
add server_test.go
#### 尝试运行测试
#### 编写最少量的代码让测试运行起来，然后检查错误输出
#### 编写足够的代码让它通过
### 重构