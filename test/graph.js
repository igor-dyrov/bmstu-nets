class Graph {
	constructor(arrOfEdeges) {
		this._vertices = arrOfEdeges;
		this._adjacency_matrix = [];
		this._vertices.forEach((vertex) => this._adjacency_matrix.push([]));
		this._vertices.forEach((vertex, index) => this._vertices.forEach((v) => this._adjacency_matrix[index].push(false)));
	}

	addEdge(from, to) {
		if (this._vertices.length < from || this._vertices.length < to) {
			return;
		} 
		this._adjacency_matrix[from][to] = true;
		this._adjacency_matrix[to][from] = true;
	}

	getNextVertices(number) {
		if (this.length < number) {
			return;
		}
		const result = [];
		this._adjacency_matrix[number].forEach((b, i) => {
			if (b) {
				result.push(i);
			}
		}) 
		return result;
	}

	bfs(number) {
		const _used = {};
		const queue = [];
		queue.push(number);
		while (queue.length) {
			const tempVertex = queue.shift();
			_used[tempVertex] = true;
			this.getNextVertices(tempVertex).forEach((v) => {
				if (!_used[v]) {
					_used[v] = true;
					queue.push(v);
				}
			});
			console.log(tempVertex);
		}
	}

}

const graph = new Graph([0, 1, 2, 3]);
graph.addEdge(1, 3);
graph.addEdge(2, 3);
graph.addEdge(2, 0);
graph.addEdge(0, 3);

graph.bfs(1);



