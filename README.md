# queueInGo
Creating a queue in golang

Introduction
At the end of this implementation, I discovered a lot of things about golang and Data Structures
Golang is a realy soft lengage and very simple to learn and understund, I got it using his own document named: "An Introduction to Programming in Go"

Objective
My first objective was to resolve a exercice that a profissional programmer gave to me
this one consists in: 
To implement a queue that has the folloing functions:
  Enqueue -  to add an iten to the queue
  Dequeue - to remove an iten from the queue
  Rear - to get the last iten from the queue
  Front - to get the first iten from the queue
  size - to get the lenght of the queue
  
 *How does the queue works?*
 
 1º step
 first you need to creat a 'new queue', this consists in just to informate
 his capacity(how many itens it will can storage) to do it we uses the following command:
  var queueExemple *Queue = newQueue([HERE THE CAPACITY])
    eg: var queueExemple*Queue = newQueue(100)
      to create a queue called queueExemple with 100 spaces

 2º step
 you can enqueue a element using the following command:
  (for this, lets use the queueExemple created in the first step at the line 22)
  enqueue([QUEUE_NAME], [ITEM_THAT_WILL_BE_ENQUEUED])
    eg: enqueue(queueExemple,1)
      to enqueue the integer 1 to the end of the queue, called queueExemple that we created in the first step.
      
 3º step
 you can dequeue a element using the following command:
  (lets also use the queueExemple to do it)
  dequeue([QUEUE_NAME])
    eg: dequeue(queueExemple)
      to dequeue the first item of de queueExemple
      
 4º step
 You can get the front of of the queue using the following command:
  getFront([QUEUE_NAME])
    eg: getFront(queueExemple)
    
 to get the rear of the queue:
  getRear([QUEUE_NAME])
    eg: getRear(queueExemple)
    
 and to get the size of the queue:
  getSize([QUEUE_NAME])
    eg: getSize(queueExemple)
  

Bibliography
https://www.golang-book.com/books/intro
https://www.geeksforgeeks.org/queue-set-1introduction-and-array-implementation/

Thanks Jonatas!

  
  
  
  
