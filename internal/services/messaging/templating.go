package messaging

import (
	"codclbot/internal/domain"
	"fmt"
	"strings"
)

// FormatAboutService — описание After-AI сервиса.
func FormatAboutService() string {
	var b strings.Builder

	b.WriteString("*After-AI разработка* 🚀\n")
	b.WriteString("*Подробности: https://codcl.com, https://t.me/aftercoding, https://medium.com/@codcl* \n\n")

	b.WriteString("Мы дорабатываем прототипы, созданные с помощью *AI* и *No-Code/Low-Code* решений:\n")
	b.WriteString("• Vercel v0, Bolt, Lovable, Replit и любых иных No/Low-code систем\n")
	b.WriteString("• Bubble, Tilda, Webflow\n\n")

	b.WriteString("Мы помогаем превратить ваш прототип в надёжный продакшн, усиливая:\n")
	b.WriteString("• backend-архитектуру\n")
	b.WriteString("• структуру базы данных\n")
	b.WriteString("• инфраструктуру и защиту\n")

	return b.String()
}

// FormatAboutTeam — короткое описание команды.
func FormatAboutTeam() string {
	var b strings.Builder

	b.WriteString("*Кто мы?*\n\n")
	b.WriteString("• Стек: Go, Python, TS/JS, Docker, Kubernetes, Cloud\n")
	b.WriteString("• Работаем в режиме *After-AI*: быстрое укрепление прототипов и доведение до продакшна\n\n")
	b.WriteString("• Тимлид: *Александр Шаман* — 20+ лет в архитектуре и разработке, AI, Web3, DevSecOps\n")

	b.WriteString("*Как мы работаем:*\n")
	b.WriteString("1. Анализ вашего прототипа или No-Code/Low-Code решения\n")
	b.WriteString("2. Формирование технического плана работ\n")
	b.WriteString("3. Оценка стоимости и сроков\n")
	b.WriteString("4. Этапная реализация с отчётностью и прозрачностью\n")

	return b.String()
}

// FormatCalculationResult — форматирует ответ калькулятора.
func FormatCalculationResult(res domain.CalculationResult, functions int) string {
	var b strings.Builder

	b.WriteString("*Черновая оценка стоимости доработки* 💶\n\n")

	b.WriteString(fmt.Sprintf("• Фронтенд: *%d €*\n", res.FrontendCost))
	b.WriteString(fmt.Sprintf("• Backend (%d функций): *%d €*\n", functions, res.BackendCost))
	b.WriteString(fmt.Sprintf("• База данных: *%d €*\n", res.DatabaseCost))
	b.WriteString(fmt.Sprintf("• Инфраструктура: *%d €*\n", res.InfraCost))
	b.WriteString("────────────────────\n")
	b.WriteString(fmt.Sprintf("*Итого: %d €*\n\n", res.Total))

	b.WriteString("_Это предварительная оценка._\n")
	b.WriteString("Стоимость зависит от конкретных интеграций, требований по безопасности, инфраструктурных ограничений и SLA.\n\n")

	b.WriteString("Чтобы получить точную смету — отправьте ссылку на ваш прототип (Vercel v0, Bolt, Bubble, Tilda, Webflow и др.).")

	return b.String()
}

// FormatWelcome — приветствие /start.
func FormatWelcome() string {
	var b strings.Builder

	b.WriteString("*Привет!* Это бот-калькулятор *After-AI разработки* 🚀\n\n")
	b.WriteString("Я помогу вам:\n")
	b.WriteString("• понять, как работает подход After-AI\n")
	b.WriteString("• познакомиться с командой\n")
	b.WriteString("• рассчитать стоимость доработки вашего прототипа\n\n")
	b.WriteString("Выберите действие с помощью кнопок ниже 👇")

	return b.String()
}
