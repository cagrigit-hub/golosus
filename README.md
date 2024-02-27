# Golosus

Golosus is a powerful Go package designed to streamline the creation of folder templates for web development projects. This package simplifies the setup process by generating a structured template with essential files for a modern web application. Golosus includes support for:

- **Templ:** Utilize the power of Go templating to customize your project structure with templ package.
- **HTMX:** Seamless integration for dynamic and interactive web pages.
- **Typescript:** Jumpstart your frontend development with a pre-configured TypeScript setup.
- **Echo:** Leverage the Echo web framework for robust backend development.

## Features

- 🌐 **Web Essentials:** Automatic creation of folders and files for a well-organized web project.
- 🚀 **Fast Setup:** Save time and effort by using Golosus to kickstart your development environment.
- 🧰 **Customization:** Tailor the template to your specific project requirements using Go templating.

## Usage (it is up to date but will be changed)

0. You need to have go & node & templ cli installed on your computer.
   I assumed you already have go & node to get templ -> [Check](https://github.com/a-h/templ)

1. Install latest Golosus from releases and add it to your bin. (this step will be changed - go get / install will be available)

2. Run Golosus to create your project template:

   ```bash
   golosus -name="YOUR PROJECT NAME" -github="YOUR GITHUB NICKNAME"
   ```

3. Install dependencies by make

   ```bash
   make init
   ```

4. or Install dependencies by hand
   ```bash
   templ generate
   go mod tidy
   npm install --prefix ./typescript
   ```
5. Start using! Check localhost:3000/example!

## Future Changes

Exciting updates are planned for Golosus! In the future, we are gearing up to introduce React support as the frontend alongside HTMX. Here's what you can expect:

- **React Integration:** Golosus will support React for the frontend, working seamlessly with HTMX. React will emit events, and HTMX will catch and reflect these changes, offering a powerful combination for dynamic and interactive web applications.

Stay tuned for upcoming releases that will bring these enhancements to Golosus, making it even more versatile and adaptable to your web development needs!

## Contributions

Contributions and feedback are welcome! Feel free to:

- **Report Issues:** Open [issues](https://github.com/cagrigit-hub/golosus/issues) to report bugs or request features.
- **Contribute Code:** Submit [pull requests](https://github.com/cagrigit-hub/golosus/pulls) to enhance Golosus and make it even more versatile.

## License

This project is licensed under the [MIT License](LICENSE.md). See the [LICENSE](LICENSE.md) file for details.
