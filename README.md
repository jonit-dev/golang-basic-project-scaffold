# Go - Project Scaffold

For structuring a Go project, a common approach aligns with the recommendations from the Go community and involves organizing your project into specific directories that each serve a distinct purpose. Here’s a straightforward and effective structure:

1. **`/cmd`** - This directory contains the main applications for your project. Each application should have its own subdirectory with a name that matches the executable you want to build. The code in these subdirectories should be minimal, ideally just parsing command line options and starting the application processes.

2. **`/pkg`** - Contains library code that's ok to use by external applications. Other projects will import these libraries expecting them to work, so think of the `/pkg` directory as providing APIs for external users.

3. **`/internal`** - Houses private application and library code. This code isn't intended to be imported by other applications. You can structure the code under this directory similar to `/pkg`, but it's meant for internal use.

4. **`/api`** - API definitions for your application. This could include definitions like OpenAPI/Swagger specs, Protobuf files, JSON schema files, etc.

5. **`/configs`** - Configuration file templates or default configs.

6. **`/static`** - Any static files that need to be served. This might include files like templates or web assets.

7. **`/scripts`** - Scripts to perform build, install, analysis, etc.

8. **`/build`** - Packaging and Continuous Integration.

9. **`/deployments`** - System and application configurations for deployment systems.

10. **`/test`** - Additional external test apps and test data. It can be useful to have a space dedicated to complex test applications.

11. **`/docs`** - Design and user documents, project documentation.

12. **`/tools`** - Supporting tools for this project. Go source files in this directory are compiled and placed in `$GOPATH/bin`.

13. **`/examples`** - Examples for your applications and/or public libraries.

14. **`/third_party`** - External helper tools, forked code, and other 3rd party utilities (e.g., Swagger UI).

15. **`/vendor`** - Application dependencies (managed manually or by your package manager like Go Modules).

**Best Practices:**

- Use Go Modules for dependency management.
- Keep your top-level package lean; it should only contain minimal code that ties together modules/components from `/internal` and `/pkg`.
- Make good use of the `internal` package to ensure encapsulation.
- Think about backward compatibility if your `/pkg` directory is going to be imported by other applications.

This structure is scalable and helps in maintaining a clear separation of concerns, enhancing understandability and maintainability of the code.
