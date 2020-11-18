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
  
  ``` git clone https://github.com/paulorubensbarbosa/queueInGo ```
  
  **Now enter in the directory:**
    ```cd $GOPATH/src/queueInGo```
  **And instal using:**
    ```go install```
  
## How to use:
  
  **First inside of your golang file, you need to import the Package using:**
  
  ```import "queueInGo"```
  
  **Now you can use all the functions inside of "queueInGo"**
  
## Functions in "queueInGo"

  **For to create a new queue you need to define his name ``[QUEUE_NAME]`` and his capacity ``[QUEUE_CAPACITY]`` using the line command:
  
  ```var [QUEUE_NAME] = queueInGo.NewQueue([QUEUE_CAPACITY])```
  
  eg
  
  ```var queueExemple = queueInGo.NewQueue(10)```
  
  **To enqueue**
  
  ```queueInGo.Enqueue([QUEUE_NAME], [ITEM_TO_BE_ENQUEUED])```
  
  **To dequeue**
  
  ```queueInGo.Dequeue([QUEUE_NAME])```
  
  **To get the first item**
  
  ```queueInGo.GetFront([QUEUE_NAME])```
  
  **To get the last item**
  
  ```queueInGo.GetRear([QUEUE_NAME])```
  
  **To get the size of the queue**
  
  ```queueInGo.GetSize([QUEUE_NAME])```

## Bibliography

- https://www.golang-book.com/books/intro

- https://www.geeksforgeeks.org/queue-set-1introduction-and-array-implementation/
