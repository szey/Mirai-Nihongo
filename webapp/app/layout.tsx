import './globals.css';
import type { Metadata } from 'next';

export const metadata: Metadata = {
  title: 'Mirai Nihongo',
  description: 'Web-first Japanese learning app for beginners',
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  );
}
