# Interview-prep
1. Describe M:N scheduler in golang.
     - Scheduler maps M goroutines onto N os threads. Each os thread has processor and local queue.
     - If queues are not full, it stacks there goroutines, If queues are full, it stacks goroutines on global queue.
     - If an os thread finds its local queue empty (and there are no goroutines in the global queue), it will attempt work stealing.
     - Work stealing is when a os thread takes (or "steals") goroutines from the local queue of another P to balance the workload and ensure efficient CPU utilization.
     - This allows the system to balance the load dynamically, reducing the risk of under-utilized threads when some P instances are busy while others are idle.
2. Describe how garbage collector works.
     - It starts with all objects initialy white
     - Each object has list of refecences
     - Root becomes gray
     - Algorithm traverses through objects' references marking objects gray. The object becomes black if all references become gray.
     - It continues until there are no more gray objects
  
