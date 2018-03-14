# dumb-queue
A horrible queue implementation

## Purpose
As an interview exercise for a python position I was asked to implement a queue using two stacks. My initial take on the task made puts fast( O(1) ) and gets slow(er) ( O(length of the queue) ). 

This project is to explore various behaviors with differnet design choices. And to push my Go skills.

## Experiments
 1. Expensive Get
 1. Expensive Put
 1. Queue with a put/get mode, switch modes as often as needed

In theory, all of the design choices should actually run an amortized O(1), but there should be measurable performance differences between the mode-queue and the fixed queues.


