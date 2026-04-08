import Link from 'next/link';

export default function HomePage() {
  return (
    <main className="page">
      <div className="container hero">
        <section>
          <span className="eyebrow">Web-first Japanese learning</span>
          <h1>Mirai Nihongo helps beginners start Japanese with a clear path.</h1>
          <p className="muted">
            The first goal is simple: turn Pre-N5 into a real, structured beginner experience with kana,
            starter reading, tiny vocabulary sets, quiz loops, and visible progress.
          </p>
          <div className="nav">
            <Link className="pill" href="/pre-n5">Open Pre-N5</Link>
            <Link className="pill" href="/status">Project Status</Link>
          </div>
        </section>
        <aside className="card">
          <span className="eyebrow">Current stage</span>
          <h2>Pre-N5 Foundation</h2>
          <p className="muted">
            A beginner-first stage focused on Hiragana, Katakana, starter reading, greetings, numbers,
            and simple quiz-based reinforcement.
          </p>
          <div className="list">
            <div><strong>Frontend</strong><br /><span className="muted">Next.js app scaffold</span></div>
            <div><strong>Backend</strong><br /><span className="muted">Go API will serve content data</span></div>
            <div><strong>Content</strong><br /><span className="muted">Stored in repository under content/pre-n5/v1</span></div>
          </div>
        </aside>
      </div>
    </main>
  );
}
