import { cn } from "@/lib/utils";
import { useState } from "react";
import { Button } from "../ui/button";
import { Github, Menu, Search } from "lucide-react";
import { Sheet, SheetContent, SheetTrigger } from "../ui/sheet";

export function Navbar() {
  const [isOpen, setIsOpen] = useState(false);

  const pathname = "/";

  const navigation = [
    { name: "Home", href: "/" },
    { name: "Docs", href: "/#" },
    { name: "Plugins", href: "/#" },
    { name: "About", href: "/#" },
  ];

  return (
    <header className="sticky top-0 z-50 w-full border-b bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60">
      <div className="container flex h-16 items-center justify-between">
        <div className="flex items-center space-x-4">
          <div className="flex items-center space-x-2 hover:opacity-80 transition-opacity">
            <div className="h-8 w-8 rounded bg-primary flex items-center justify-center">
              <span className="text-primary-foreground font-bold text-sm">
                H
              </span>
            </div>
            <span className="font-bold text-xl">httpkit</span>
          </div>
        </div>

        {/* Desktop Navigation */}
        <nav className="hidden md:flex items-center space-x-1">
          {navigation.map((item) => (
            <div
              key={item.name}
              className={cn(
                "px-3 py-2 text-sm font-medium transition-colors rounded-md",
                "hover:text-primary hover:bg-accent/50",
                pathname === item.href ||
                  (item.href !== "/" && pathname.startsWith(item.href))
                  ? "text-primary bg-accent/50"
                  : "text-muted-foreground"
              )}
            >
              {item.name}
            </div>
          ))}
        </nav>

        <div className="flex items-center space-x-2">
          <Button
            variant="ghost"
            size="sm"
            className="hidden sm:flex text-muted-foreground hover:text-foreground"
          >
            <Search className="h-4 w-4 mr-2" />
            Search
            <kbd className="ml-2 pointer-events-none inline-flex h-5 select-none items-center gap-1 rounded border bg-muted px-1.5 font-mono text-[10px] font-medium text-muted-foreground opacity-100">
              <span className="text-xs">⌘</span>K
            </kbd>
          </Button>
          <Button variant="ghost" size="sm" asChild>
            <a href="https://github.com/hermangoncalves/httpkit" target="_blank">
              <Github className="h-4 w-4" />
              <span className="sr-only">GitHub</span>
            </a>
          </Button>
          {/* <ThemeToggle /> */}

          {/* Mobile Menu */}
          <Sheet open={isOpen} onOpenChange={setIsOpen}>
            <SheetTrigger asChild className="md:hidden">
              <Button variant="ghost" size="sm">
                <Menu className="h-4 w-4" />
                <span className="sr-only">Toggle menu</span>
              </Button>
            </SheetTrigger>
            <SheetContent side="right" className="w-[300px] sm:w-[400px]">
              <div className="flex flex-col space-y-4 mt-8">
                <div className="flex items-center justify-between mb-4">
                  <span className="font-semibold">Navigation</span>
                </div>
                {navigation.map((item) => (
                  <a
                    key={item.name}
                    href={item.href}
                    className={cn(
                      "text-sm font-medium transition-colors py-2 px-3 rounded-md",
                      "hover:text-primary hover:bg-accent/50",
                      pathname === item.href ||
                        (item.href !== "/" && pathname.startsWith(item.href))
                        ? "text-primary bg-accent/50"
                        : "text-muted-foreground"
                    )}
                    onClick={() => setIsOpen(false)}
                  >
                    {item.name}
                  </a>
                ))}
                <div className="border-t pt-4 mt-4 space-y-2">
                  <Button variant="ghost" className="w-full justify-start">
                    <Search className="mr-2 h-4 w-4" />
                    Search Documentation
                  </Button>
                  <Button
                    variant="ghost"
                    className="w-full justify-start"
                    asChild
                  >
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
            </SheetContent>
          </Sheet>
        </div>
      </div>
    </header>
  );
}
