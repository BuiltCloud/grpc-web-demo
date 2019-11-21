using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading;
using System.Threading.Tasks;
using Grpc.Core;
using Microsoft.Extensions.Logging;

namespace GrpcService1
{
    public class GreeterService : Greeter.GreeterBase
    {
        private readonly ILogger<GreeterService> _logger;
        public GreeterService(ILogger<GreeterService> logger)
        {
            _logger = logger;
        }

        public override Task<HelloReply> SayHello(HelloRequest request, ServerCallContext context)
        {
            return Task.FromResult(new HelloReply
            {
                Message = "Hello " + request.Name
            });
        }

        public override Task<HelloReply> SayHelloAfterDelay(HelloRequest request, ServerCallContext context)
        {
            Thread.Sleep(2000);
            return Task.FromResult(new HelloReply
            {
                Message = "SayHelloAfterDelay " + request.Name
            });
        }
        

        public override Task SayRepeatHello(RepeatHelloRequest request,IServerStreamWriter<HelloReply> responseStream, ServerCallContext context)
        {
            return Task.Run(async()=> {
                for (var i = 0; i < request.Count; i++) {
                    await responseStream.WriteAsync(new HelloReply
                    {
                        Message = "Hello " + request.Name + i
                    });
                }
                
            });
        }
    }
}
