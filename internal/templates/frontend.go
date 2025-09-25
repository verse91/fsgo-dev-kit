package templates

// HeroComponent returns the Hero component template
func HeroComponent() string {
	return `export default function Hero() {
  return (
    <div className="hero">
      <h1>Welcome to your new project!</h1>
    </div>
  );
}`
}

// NavbarComponent returns the Navbar component template
func NavbarComponent() string {
	return `export default function Navbar() {
  return (
    <nav className="navbar">
      <div>Your App</div>
    </nav>
  );
}`
}

// TypographyComponent returns the Typography component template
func TypographyComponent() string {
	return `export const Typography = {
  h1: ({ children, ...props }: React.HTMLAttributes<HTMLHeadingElement>) => (
    <h1 className="text-4xl font-bold" {...props}>{children}</h1>
  ),
  h2: ({ children, ...props }: React.HTMLAttributes<HTMLHeadingElement>) => (
    <h2 className="text-3xl font-semibold" {...props}>{children}</h2>
  ),
  p: ({ children, ...props }: React.HTMLAttributes<HTMLParagraphElement>) => (
    <p className="text-base" {...props}>{children}</p>
  ),
};`
}

// SignInPage returns the Sign In page template
func SignInPage() string {
	return `export default function SignIn() {
  return (
    <div className="min-h-screen flex items-center justify-center">
      <div className="max-w-md w-full space-y-8">
        <h2 className="text-center text-3xl font-extrabold text-gray-900">
          Sign in to your account
        </h2>
      </div>
    </div>
  );
}`
}

// AuthCallbackPage returns the Auth Callback page template
func AuthCallbackPage() string {
	return `export default function AuthCallback() {
  return (
    <div className="min-h-screen flex items-center justify-center">
      <div className="text-center">
        <h2 className="text-2xl font-bold">Processing authentication...</h2>
      </div>
    </div>
  );
}`
}

// FrontendEnvFile returns the frontend environment file template
func FrontendEnvFile() string {
	return "# Environment variables\nNEXT_PUBLIC_API_URL=http://localhost:8080\n"
}

// FrontendEnvExampleFile returns the frontend environment example file template
func FrontendEnvExampleFile() string {
	return "# Environment variables example\nNEXT_PUBLIC_API_URL=http://localhost:8080\n"
}