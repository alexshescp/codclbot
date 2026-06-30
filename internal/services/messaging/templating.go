package messaging

import (
	"codclbot/internal/domain"
	"fmt"
	"strings"
)

// FormatAboutService — описание EURO SMM сервиса.
func FormatAboutService() string {
	var b strings.Builder

	b.WriteString("*EURO SMM — премиальная архитектура соцсетей* 🚀\n")
	b.WriteString("Мы создаём визуальные экосистемы, которые усиливают бренд, растят органику, SEO и доверие аудитории. Это не просто соцсети — это инфраструктура цифрового присутствия.\n\n")

	b.WriteString("*Ключевые метрики*\n")
	b.WriteString("• 24ч — среднее время запуска\n")
	b.WriteString("• 99% — удовлетворённость клиентов\n")
	b.WriteString("• 3× — ускорение роста площадок\n")
	b.WriteString("• 50+ — подключённых платформ\n\n")

	b.WriteString("*Основные направления*\n")
	b.WriteString("• Премиальные пакеты: Europack, Euromaxpack, International Maxpack, Landing Chain Generator\n")
	b.WriteString("• SEO-упаковка под любую конкурентность\n")
	b.WriteString("• Контент-маркетинг для поиска и соцсетей\n")
	b.WriteString("• Встроенный калькулятор стоимости и консультация в один клик\n\n")

	b.WriteString("Чтобы изучить полный каталог услуг, нажмите кнопку *«Каталог услуг»* в меню. Остальной функционал бота — калькулятор пакета и заявка на консультацию — всегда доступен.")

	return b.String()
}

// FormatServiceCatalog — полный каталог услуг и пакетов.
func FormatServiceCatalog() string {
	var b strings.Builder

	b.WriteString("*Каталог услуг EURO SMM*\n")
	b.WriteString("Премиальная визуальная архитектура, связанная с SEO, контент-цепочками и расширением бренда на всех ключевых площадках.\n\n")

	b.WriteString("*Преимущества*\n")
	b.WriteString("• Запуск, оптимизация и аналитика — в одном цикле\n")
	b.WriteString("• Визуальная система: айдентика, motion, сторителлинг\n")
	b.WriteString("• SEO + соцсети: рост под алгоритмы поиска и платформ\n")
	b.WriteString("• Время запуска: от 24 часов\n\n")

	b.WriteString("*Метрики*\n")
	b.WriteString("• 24ч — запуск\n")
	b.WriteString("• 99% — довольных клиентов\n")
	b.WriteString("• 3× — рост показателей\n")
	b.WriteString("• 50+ — площадок\n\n")

	b.WriteString("*Премиальные пакеты*\n")
	b.WriteString("• Europack — европейское присутствие (€299, 2d)\n")
	b.WriteString("• Euromaxpack — полное европейское доминирование (€599, 4d)\n")
	b.WriteString("• International Maxpack — глобальный охват (€999, 7d)\n")
	b.WriteString("• Landing Chain Generator — система конверсии из лендингов (€499, 5d)\n\n")

	b.WriteString("*Детали пакетов*\n")
	b.WriteString("• Europack: 20 единиц контента, 6 площадок, базовое SEO, бренд-кит, аналитика, поддержка 24/7\n")
	b.WriteString("• Euromaxpack (топ-пакет): 50 единиц контента, 8 площадок, расширенная стратегия, полный SEO, премиальный бренд-кит, расширенная аналитика, анализ конкурентов, ежемесячные отчёты\n")
	b.WriteString("• International Maxpack: 100 единиц контента, 20+ площадок, Enterprise-SEO, мастер-бренд-кит, ИИ-аналитика, выделенный менеджер, еженедельные отчёты, кастомные интеграции\n")
	b.WriteString("• Landing Chain Generator: 5 взаимосвязанных лендингов, A/B-тесты, аналитика, SEO-архитектура, оптимизация конверсии\n\n")

	b.WriteString("*SEO-упаковка*\n")
	b.WriteString("• Низкая конкуренция (€300): кластеризация, OG, Schema, онпейдж, скорость, локальные ключи\n")
	b.WriteString("• Средняя (€2000): Entity SEO, топикальные силосы, белые ссылки, обновление ключевых страниц, интент-оптимизация\n")
	b.WriteString("• Высокая (€9000): техническое SEO, контент-маркетинг, полный тех-аудит, расширенные сниппеты, authoritative-ссылки\n\n")

	b.WriteString("*Контент-маркетинг*\n")
	b.WriteString("• 10 статей/мес + кросспостинг (50 публикаций) — €1199\n")
	b.WriteString("• 20 статей/мес + кросспостинг — €1999\n")
	b.WriteString("• YouTube-комплект: 60 минут видео/мес — €5000\n")
	b.WriteString("• Вирусные видео: от €10000 + гарантированный буст\n")
	b.WriteString("• Лонгриды и дистрибуция на LinkedIn/Medium/X/Reddit/Substack\n\n")

	b.WriteString("*Философия*\n")
	b.WriteString("• Оркестровка каналов: единая визуальная и смысловая система\n")
	b.WriteString("• Дизайн дохода: архитектура конверсий под поиск и соцсети\n")
	b.WriteString("• Креативная инженерия: контент-киты и дизайн-системы\n")
	b.WriteString("• Рост: evergreen-контент + платное усиление и комьюнити-эффекты\n\n")

	b.WriteString("*Истории успеха*\n")
	b.WriteString("• AI-платформа: +300% роста за 3 месяца на 12 площадках\n")
	b.WriteString("• Fashion-бренд: +500% вовлечения за 90 дней\n")
	b.WriteString("• SaaS B2B: 200 лидов в месяц с LinkedIn-экосистемы\n")
	b.WriteString("• Медиа-сеть: 1M+ охвата ежемесячно\n\n")

	b.WriteString("*Калькулятор услуг*\n")
	b.WriteString("Соберите ваш пакет: бренд-база, стратегический охват, количество кампаний, контент-система и каналы — итоговая смета формируется мгновенно.\n\n")

	b.WriteString("*Как начать*\n")
	b.WriteString("• Выберите пакет или откройте калькулятор\n")
	b.WriteString("• Оставьте контакты — ответ в течение 24 часов\n")
	b.WriteString("• Укажите цели и ключевые площадки (X, LinkedIn, Instagram, Dribbble и др.)\n\n")

	b.WriteString("Пример заявки: Имя / Email / компания / выбранный пакет / цели проекта.\n")
	b.WriteString("_Нажимая кнопку, вы подтверждаете согласие с политикой конфиденциальности._")

	return b.String()
}

// FormatAboutTeam — короткое описание команды.
func FormatAboutTeam() string {
	var b strings.Builder

	b.WriteString("*Наша философия*\n")
	b.WriteString("• Оркестровка каналов — единый сторителлинг в платных, собственных и earned-медиа\n")
	b.WriteString("• Дизайн дохода — маршруты конверсии под поиск, соцсети и комьюнити\n")
	b.WriteString("• Креативные процессы — модульные контент-киты, дизайн-системы и регламенты\n")
	b.WriteString("• Инженерия роста — контент-цепочки, SEO-архитектура, evergreen-механики\n")
	b.WriteString("• Сообщество — визуальная идентичность на всех точках контакта\n\n")

	b.WriteString("*Оценка зрелости*\n")
	b.WriteString("• Стратегия — 92% (ICP, позиционирование, value ladder)\n")
	b.WriteString("• Дизайн — 88% (бренд-OS, motion, UI-киты)\n")
	b.WriteString("• Контент — 94% (evergreen + rapid content)\n")
	b.WriteString("• SEO — 90% (entity-first, техника, OG)\n")
	b.WriteString("• Рост — 96% (комьюнити + платное усиление)\n\n")

	b.WriteString("*Кейсы*\n")
	b.WriteString("• AI-платформа: +300% роста за 3 месяца\n")
	b.WriteString("• Fashion-бренд: +500% вовлечения\n")
	b.WriteString("• SaaS Enterprise: 200 лидов/мес через LinkedIn\n")
	b.WriteString("• Медиа-проект: 1M+ охвата ежемесячно\n\n")

	b.WriteString("Готовы к премиальной трансформации вашего цифрового присутствия? Оставьте заявку — ответ в течение 24 часов.")

	return b.String()
}

// FormatCalculationResult — форматирует ответ калькулятора.
func FormatCalculationResult(res domain.CalculationResult, functions int) string {
	var b strings.Builder

	b.WriteString("*Черновая смета пакета EURO SMM* 💶\n\n")

	b.WriteString(fmt.Sprintf("• Бренд-база и визуальное присутствие: *%d €*\n", res.FrontendCost))
	b.WriteString(fmt.Sprintf("• Стратегия и охват (%d кампаний): *%d €*\n", functions, res.BackendCost))
	b.WriteString(fmt.Sprintf("• Контент-система и производство: *%d €*\n", res.DatabaseCost))
	b.WriteString(fmt.Sprintf("• Архитектура каналов и дистрибуция: *%d €*\n", res.InfraCost))
	b.WriteString("────────────────────\n")
	b.WriteString(fmt.Sprintf("*Итого: %d €*\n\n", res.Total))

	b.WriteString("_Это предварительная оценка._\n")
	b.WriteString("Финальная стоимость зависит от числа площадок, уровня конкуренции, объёма SEO-работ и глубины контент-маркетинга.\n\n")

	b.WriteString("Чтобы зафиксировать пакет — отправьте цели проекта и ключевые площадки (X, LinkedIn, Instagram, Dribbble и др.).")

	return b.String()
}

// FormatWelcome — приветствие /start.
func FormatWelcome() string {
	var b strings.Builder

	b.WriteString("*Добро пожаловать!* Это EURO SMM — премиальная архитектура соцсетей.\n\n")
	b.WriteString("Я помогу вам:\n")
	b.WriteString("• изучить *каталог услуг*\n")
	b.WriteString("• выбрать пакет Europack / Euromaxpack / International Maxpack\n")
	b.WriteString("• получить SEO-упаковку под конкурентность\n")
	b.WriteString("• собрать контент-маркетинг под поиск и соцсети\n")
	b.WriteString("• посмотреть философию и кейсы команды\n\n")
	b.WriteString("Выберите действие с помощью кнопок ниже 👇")

	return b.String()
}
