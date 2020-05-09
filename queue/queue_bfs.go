package queue

/*
	队列和BFS
	广度优先搜索（BFS）的一个常见应用是找出从根节点到目标节点的最短路径
*/

// 组成图的顶点
type Node struct {
	Value interface{}
}

type NodeQueue struct {
	Items []Node
}

func NewNodeQueue() *NodeQueue {
	n := new(NodeQueue)
	n.Items = []Node{}
	return n
}

func (n *NodeQueue) Enqueue(t Node) {
	n.Items = append(n.Items, t)
}

func (n *NodeQueue) Dequeue() *Node {
	if n.IsEmpty() {
		return nil
	}

	item := n.Items[0]
	n.Items = n.Items[1:len(n.Items)]
	return &item
}

func (n *NodeQueue) IsEmpty() bool {
	return len(n.Items) == 0
}

// 定义一个图的结构，图有顶点与边组成 V E
type ItemGraph struct {
	Nodes []*Node
	Edges map[Node][]*Node
}

func NewItemGraph() *ItemGraph {
	return &ItemGraph{
		Nodes: make([]*Node, 0),
		Edges: make(map[Node][]*Node),
	}
}

// 添加节点
func (g *ItemGraph) AddNode(n *Node) {
	g.Nodes = append(g.Nodes, n)
}

// 添加边
func (g *ItemGraph) AddEdge(n1, n2 *Node) {
	if g.Edges == nil {
		g.Edges = make(map[Node][]*Node)
	}

	//无向图
	g.Edges[*n1] = append(g.Edges[*n1], n2) // 设定从节点n1到n2的边
	g.Edges[*n2] = append(g.Edges[*n2], n1) // 设定从节点n2到n1的边
}

func (g *ItemGraph) Bfs(target *Node) (bool, int) {
	q := NewNodeQueue()             // 核心数据结构
	visited := make(map[*Node]bool) // 避免走回头路
	step := 0                       // 记录扩散的步骤

	n := g.Nodes[0] // 将起点加入队列
	q.Enqueue(*n)
	visited[n] = true

	for {
		if q.IsEmpty() {
			break
		}

		size := len(q.Items) //将当前队列中的所有节点向四周扩散
		for i := 0; i < size; i++ {
			node := q.Dequeue()
			if node.Value == target.Value {
				return true, step
			}
			near := g.Edges[*node]
			for i := 0; i < len(near); i++ {
				j := near[i]
				if !visited[j] {
					q.Enqueue(*j)
					visited[j] = true
				}
			}

		}

		step++ // 更新步数
	}

	return false, step
}
