import { Navbar } from "@/components/layout/navbar";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { ArrowRight, Github, Layers, Shield, Zap } from "lucide-react";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";

export default function Index() {
  return (
    <div className="min-h-screen bg-background container mx-auto">
      <Navbar />

      {/* Hero Section */}
      <section className="container py-24 md:py-32">
        <div className="mx-auto max-w-4xl text-center">
          <Badge variant="secondary" className="mb-4">
            Lightweight HTTP Toolkit
          </Badge>
          <h1 className="text-4xl md:text-6xl font-bold tracking-tight mb-6">
            Build extensible HTTP servers in{" "}
            <span className="text-primary">Go</span>
          </h1>
          <p className="text-xl text-muted-foreground mb-8 max-w-2xl mx-auto">
            A lightweight wrapper around Go's{" "}
            <code className="bg-muted px-2 py-1 rounded text-sm">net/http</code>{" "}
            package with a flexible plugin and middleware system.{" "}
            <strong>Note:</strong> httpkit is <em>not</em> an HTTP router.
          </p>
          <div className="flex flex-col sm:flex-row gap-4 justify-center">
            <Button size="lg" asChild>
              <a href="/#quick-start">
                Get Started
                <ArrowRight className="ml-2 h-4 w-4" />
              </a>
            </Button>
            <Button variant="outline" size="lg" asChild>
              <a
                href="https://github.com/hermangoncalves/httpkit"
                target="_blank"
              >
                <Github className="mr-2 h-4 w-4" />
                View on GitHub
              </a>
            </Button>
          </div>
        </div>
      </section>

      <section className="container py-16 border-t">
        <div className="mx-auto max-w-4xl">
          <h2 className="text-3xl font-bold text-center mb-12">Why httpkit?</h2>
          <div className="grid md:grid-cols-3 gap-8">
            <Card>
              <CardHeader>
                <Zap className="h-8 w-8 text-primary mb-2" />
                <CardTitle>Lightweight</CardTitle>
                <CardDescription>
                  No routing layer, just <code>net/http</code> underneath
                </CardDescription>
              </CardHeader>
            </Card>
            <Card>
              <CardHeader>
                <Layers className="h-8 w-8 text-primary mb-2" />
                <CardTitle>Flexible</CardTitle>
                <CardDescription>
                  Install plugins globally or per handler
                </CardDescription>
              </CardHeader>
            </Card>
            <Card>
              <CardHeader>
                <Shield className="h-8 w-8 text-primary mb-2" />
                <CardTitle>Extensible</CardTitle>
                <CardDescription>
                  Create custom middleware and plugins
                </CardDescription>
              </CardHeader>
            </Card>
          </div>
        </div>
      </section>

      <section className="container py-16">
        <div className="mx-auto max-w-4xl">
          <h2 className="text-3xl font-bold text-center mb-12">
            Core Features
          </h2>
          <div className="space-y-6">
            <div className="flex items-start space-x-4">
              <div className="w-2 h-2 bg-primary rounded-full mt-2"></div>
              <div>
                <h3 className="font-semibold mb-2">
                  Plugin-Based Architecture
                </h3>
                <p className="text-muted-foreground">
                  Encapsulate functionality in reusable plugins that can be
                  applied globally or per-route.
                </p>
              </div>
            </div>
            <div className="flex items-start space-x-4">
              <div className="w-2 h-2 bg-primary rounded-full mt-2"></div>
              <div>
                <h3 className="font-semibold mb-2">
                  Context-Based Request Handling
                </h3>
                <p className="text-muted-foreground">
                  Enhanced context wrapping <code>http.ResponseWriter</code> and{" "}
                  <code>*http.Request</code> with helpful utilities.
                </p>
              </div>
            </div>
            <div className="flex items-start space-x-4">
              <div className="w-2 h-2 bg-primary rounded-full mt-2"></div>
              <div>
                <h3 className="font-semibold mb-2">Middleware Chain</h3>
                <p className="text-muted-foreground">
                  Compose middleware functions to modify request/response
                  processing in a predictable order.
                </p>
              </div>
            </div>
            <div className="flex items-start space-x-4">
              <div className="w-2 h-2 bg-primary rounded-full mt-2"></div>
              <div>
                <h3 className="font-semibold mb-2">Service Sharing</h3>
                <p className="text-muted-foreground">
                  Share services and data between middleware and handlers via
                  request context.
                </p>
              </div>
            </div>
          </div>
        </div>
      </section>

      <section id="quick-start" className="container py-16 border-t">
        <div className="mx-auto max-w-4xl">
          <h2 className="text-3xl font-bold text-center mb-12">
            Quick Example
          </h2>
          <Card>
            <CardHeader>
              <CardTitle className="flex items-center justify-between">
                main.go
                <Badge variant="secondary">Go</Badge>
              </CardTitle>
            </CardHeader>
            <CardContent>
              <pre className="bg-muted p-4 rounded-lg overflow-x-auto">
                <code className="text-sm">{`// Install
go get github.com/hermangoncalves/httpkit

// main.go
package main

import (
    "github.com/hermangoncalves/httpkit"
)

func main() {
    app := httpkit.New()

    // Register global plugins
    app.RegisterPlugin(NewLoggerPlugin())

    // Add a handler
    app.Handle("/", func(ctx *httpkit.Context) {
        ctx.JSON(200, httpkit.H{"message": "Hello, World"})
    })

    app.Run(":8080")
}`}</code>
              </pre>
            </CardContent>
          </Card>
        </div>
      </section>

      <footer className="border-t py-12 mt-16">
        <div className="container">
          <div className="flex flex-col md:flex-row justify-between items-center">
            <div className="flex items-center space-x-2 mb-4 md:mb-0">
              <div className="h-6 w-6 rounded bg-primary flex items-center justify-center">
                <span className="text-primary-foreground font-bold text-xs">
                  H
                </span>
              </div>
              <span className="font-semibold">httpkit</span>
            </div>
            <div className="flex items-center space-x-6 text-sm text-muted-foreground">
              <a
                href="/docs"
                className="hover:text-foreground transition-colors"
              >
                Documentation
              </a>
              <a
                href="/plugins"
                className="hover:text-foreground transition-colors"
              >
                Plugins
              </a>
              <a
                href="https://github.com/hermangoncalves/httpkit"
                target="_blank"
                className="hover:text-foreground transition-colors flex items-center"
              >
                GitHub
                {/* <Externala className="ml-1 h-3 w-3" /> */}
              </a>
            </div>
            <p className="text-sm text-muted-foreground mt-4 md:mt-0">
              © 2025 Herman Gonçalves
            </p>
          </div>
        </div>
      </footer>
    </div>
  );
}
