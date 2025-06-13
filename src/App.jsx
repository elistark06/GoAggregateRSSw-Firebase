import { useState } from 'react'
import './App.css'
import { MantineProvider, Title, Card, Text } from '@mantine/core'

function App() {
  const [count, setCount] = useState(0)

  return (
      <MantineProvider>
          <Title>Custom RSS Feed</Title>

          <Card withBorder shadow="lg" p="lg" radius="md" w={"50%"} >
            <Text size='sm'>
              this is content. Lorem ipsum dolor sit amet consectetur adipisicing elit. Corporis similique quasi nihil ipsa veniam sed error corrupti iure expedita, nesciunt magni dolores vitae cum unde rerum praesentium quo obcaecati quod?
            </Text>
          </Card>
      </MantineProvider>
  )
}

export default App
