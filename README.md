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
     - Colors:
          - White means it's unreachable and to be garbage collected,
          -  Gray means it's reachable and to be proceseed in the future
          -  Black means it's reachable.
     - Concurrent and Non-blocking: One of the strengths of Go's garbage collector is that it runs concurrently with the program and aims to be non-blocking. 
      
3. Difference between array and slice
     - Arrays have a fixed size and are value types, making them suitable for when you have a known fixed number of elements.
     - Slices, in contrast, are dynamic, more versatile, and are reference types,
  
4. Why choose Goroutines?

     - The advantages of goroutines are clear:
          
          - Efficiency: They require less memory and resources, which allows you to run thousands, even millions, of concurrent tasks.
          
          - Simplicity: The syntax for creating goroutines is very simple, making concurrent and parallel programming more approachable.
          
          - Scalability: Since they're so lightweight, goroutines enable applications to scale more efficiently in terms of memory usage and management overhead.
