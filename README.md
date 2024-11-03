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
5. What are string literals?

     - A string literal is a string constant formed by concatenating characters. The two forms of string literal are raw and interpreted string literals

6. What are packages in a Go program?

     - Packages (pkg) are directories within your Go workspace that contain Go source files or other packages. Every function, variable, and type from your source files are stored in the linked package. Every Go source file belongs to a package,
  
7. How do we perform inheritance with Golang?

     - This is a bit of a trick question: there is no inheritance in Golang because it does not support classes.
     - However, you can mimic inheritance behavior using composition to use an existing struct object to define a starting behavior of a new object. Once the new object is created, functionality can be extended beyond the original struct.

8. Explain Go interfaces. What are they and how do they work?
     - Interfaces are a special type in Go that define a set of method signatures but do not provide implementations. Values of interface type can hold any value that implements those methods.
     - Interfaces essentially act as placeholders for methods that will have multiple implementations based on what object is using it.

9. Can you return multiple values from a function?
     - Yes. A Go function can return multiple values, each separated by commas in the return statement.
10. Reverse the order of a slice
     - Implement function reverse that takes a slice of integers and reverses the slice in place without using a temporary slice.
       ```
                 package main
          
          import "fmt"
          
          func reverse(sw []int) {
                  for a, b := 0, len(sw)-1; a < b; a, b = a+1, b-1 {
                          sw[a], sw[b] = sw[b], sw[a]
                  } 
          }
          
          func main() { 
              x := []int{3, 2, 1} 
              reverse(x)
              fmt.Println(x)
          }
     ```
11. What is gRPC?

     - gRPC is an open-source framework developed by Google that enables efficient communication between distributed systems by providing a language-agnostic, high-performance Remote Procedure Call (RPC) mechanism.

12. What are the advantages of using gRPC over traditional REST APIs?
     - gRPC offers advantages such as high performance, bi-directional streaming, support for multiple programming languages, automatic code generation, efficient data serialization using Protocol Buffers, and built-in authentication and load balancing.

13. What is Protocol Buffers?
     - Protocol Buffers is a language-agnostic data serialization format used by gRPC. It allows efficient encoding and decoding of structured data, making it faster and more compact compared to other formats like JSON or XML.

14. How does gRPC handle data serialization and deserialization?
     - gRPC uses Protocol Buffers for data serialization and deserialization. Protocol Buffers allow the definition of message types and generate code for various programming languages, making it easy to send and receive structured data between clients and servers.
15. How does gRPC handle error handling and status codes?
     - gRPC uses status codes to indicate the result of an RPC call. It provides a rich set of status codes, including standard HTTP status codes, to represent different error scenarios. Additionally, gRPC allows developers to define custom status codes for application-specific errors.

16. What are interceptors in gRPC?
     - Interceptors in gRPC are middleware components that can intercept and modify RPC messages both on the client and server side. They enable cross-cutting concerns like logging, authentication, or monitoring to be implemented in a reusable manner.

17. What is the difference between gRPC and GraphQL?
     - gRPC is a high-performance RPC framework for efficient communication between distributed systems. GraphQL, on the other hand, is a query language and runtime for APIs that enables clients to request specific data and shape the response according to their needs.

18. Explain the concept of the gRPC gateway.
     - gRPC gateway is a tool that generates a reverse proxy server, allowing clients to access gRPC services using traditional HTTP/JSON-based APIs. It enables interoperability between gRPC and RESTful services.




