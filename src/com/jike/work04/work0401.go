package main
import (
    "fmt"
)

var m int
var n int

func solve(board [][]byte)  {
    if len(board) < 2 || len(board[0]) < 2 {
        return
    }
    m = len(board)
    n = len(board[0])
    //对边界数据进行遍历，查找与边界相连的数据,并把值临时置为A
    for j := 0; j < n; j++ {
        if board[0][j] == 'O'  {
            dfs(board,0,j)
        }
        if board[m-1][j] == 'O'  {
            dfs(board,m-1,j)
        }
    }
    for i := 0; i < m; i++ {
        if board[i][0] == 'O'  {
            dfs(board,i,0)
        }
        if board[i][n-1] == 'O'  {
            dfs(board,i,n-1)
        }
    }
    fmt.Println(board);
    //遍历数据，将等于A的改回O，O的数据改为X
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] == 'O' {
                board[i][j] = 'X'
            }
            if board[i][j] == 'A' {
                board[i][j] = 'O'
            }
        }
    }
}

func dfs(board [][]byte,x int,y int)  {
    if x < 0 || x >= m || y < 0 || y >= n || board[x][y] != 'O' {
        return
    }

    //按上左下右顺序遍历，直到边界或者值不为O退出遍历
    board[x][y] = 'A'
    dfs(board, x - 1, y)
    dfs(board, x, y + 1)
    dfs(board, x + 1, y)
    dfs(board, x, y - 1)
}


func main(){
    aInput := [][]byte{{'X','X','X','X'},{'X','O','O','X'},{'X','X','O','X'},{'X','O','X','X'}}
    //aInput := [][]byte{{'O','X','X','O','X'},{'X','O','O','X','O'},{'X','O','X','O','X'},{'O','X','O','O','O'},{'X','X','O','X','O'}}
    solve(aInput)
    fmt.Println(aInput);
}



