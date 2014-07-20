require "resque"

class Hello
  @queue = :hello
end

100.times { Resque.enqueue(Hello) }
