type Args struct{ A, B int }
type Reply int

srv := rpc2.NewServer()
srv.Handle("add", func(client *rpc2.Client, args *Args, reply *Reply) error {
        // Reversed call (server to client)
        var rep Reply
        client.Call("mult", Args{2, 3}, &rep)
        fmt.Println("mult result:", rep)

        *reply = Reply(args.A + args.B)
        return nil
})

lis, _ := net.Listen("tcp", "127.0.0.1:5000")
srv.Accept(lis)