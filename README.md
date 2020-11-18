# queueInGo
## Creating a queue in golang

## Objective
My first objective was to resolve a exercice that a profissional programmer gave to me
this one consists in: 

- To implement a queue that has the folloing functions:  
- Enqueue -  to add an iten to the queue 
- Dequeue - to remove an iten from the queue
- Rear - to get the last iten from the queue
- Front - to get the first iten from the queue
- size - to get the lenght of the queue

## Installation

  **First in your Bash or gitBash you need to go to the correct directory, for this use the following command:**
  
  ``` cd $GOPATH/src ```
  
  **Then clone this repository using:**
  
  ``` git clone https://github.com/paulorubensbarbosa/queue ```
  
  **Now enter in the directory:**
    ```cd $GOPATH/src/queue```
  **And instal using:**
    ```go install```
  
## How to use:
  
  **First inside of your golang file, you need to import the Package using:**
  
  ```import "queue"```
  
  **Now you can use all the functions inside of "queueInGo"**
  
## Functions in "queueInGo"

  **For to create a new queue you need to define his name ``[QUEUE_NAME]`` and his capacity ``[QUEUE_CAPACITY]`` using the line command:**
  
  ```var [QUEUE_NAME] = queue.NewQueue([QUEUE_CAPACITY])```
  
  eg
  
  ```var queueExemple = queue.NewQueue(10)```
  
  **All the following Functions, just will works if you already created a queue, using the first function**
  
  **To enqueue**
  
  ```queue.Enqueue([QUEUE_NAME], [ITEM_TO_BE_ENQUEUED])```
  
  **To dequeue**
  
  ```queue.Dequeue([QUEUE_NAME])```
  
  **To get the first item**
  
  ```queue.GetFront([QUEUE_NAME])```
  
  **To get the last item**
  
  ```queue.GetRear([QUEUE_NAME])```
  
  **To get the size of the queue**
  
  ```queue.GetSize([QUEUE_NAME])```

## Bibliography

- https://www.golang-book.com/books/intro

- https://www.geeksforgeeks.org/queue-set-1introduction-and-array-implementation/

## Contact and Feed back:

- Email: paulorubensbarbosa@outlook.com
