

# Morphia Bot

Morphia is a Discord bot that allows users to **create, remove, and edit their own custom roles** directly from Discord.
It’s perfect for communities where users want personalized role names, colors, and visibility without giving full admin privileges.

---

## Features

* **Create your own role**
  Users can create a new role with a custom name and color.

* **Edit your role**
  Change your role’s name or color anytime.

* **Remove your role**
  Delete your custom role whenever you want.

* **Safe Permissions**
  Users can manage only their own roles; no admin privileges are required.

---

## Commands

* `/createrole <name> <color>` – Create a new role
* `/editrole <name> <color>` – Update your role
* `/deleterole` – Remove your role

> Role colors can be provided as hex values (e.g., `0xFF00AA`).

---

## Setup

1. Clone the repository:

```bash
git clone https://github.com/mertensblackfyre/morphia
cd morphia
```

2. Install dependencies:

```bash
go mod tidy
```


3. Run the bot:

```bash
go run .
```

---

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.


